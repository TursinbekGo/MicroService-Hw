package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	"github.com/asadbek/app/dict"
)

type server struct {
	*dict.UnimplementedTranslateServer
}

func (s *server) Dictionary(ctx context.Context, req *dict.DictionaryRequest) (*dict.DictionaryResponse, error) {

	var fruits map[string]string = map[string]string{
		"apple":  "olma",
		"orange": "apelsin",
		"cherry": "olcha",
	}

	value := fruits[req.Key]

	return &dict.DictionaryResponse{
		Value: value,
	}, nil
}

func (s *server) Add(ctx context.Context, req *dict.AddRequest) (*dict.AddResponse, error) {

	result := req.Key + req.Key1

	return &dict.AddResponse{
		Result: result,
	}, nil
}

func (s *server) GetCurrency(ctx context.Context, req *dict.CurrencyRequest) (*dict.Result, error) {
	var res []*dict.CurrencyResponse = []*dict.CurrencyResponse{}
	var result []*dict.CurrencyResponse = []*dict.CurrencyResponse{}

	resp, err := http.Get("https://cbu.uz/oz/arkhiv-kursov-valyut/json/")

	if err != nil {
		log.Fatal(err)
	}
	// defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &res)

	for _, val := range res {
		if strings.Contains(val.Ccy, req.Ccy) {
			result = append(result, val)
		}
	}

	// fmt.Println(res)

	return &dict.Result{
		Infos: result,
	}, nil
}
func (s *server) Converter(ctx context.Context, req *dict.MoneyConverterRequest) (*dict.MoneyConverterResponse, error) {
	var res []*dict.MoneyConverterResponse = []*dict.MoneyConverterResponse{}
	resp, err := http.Get("https://cbu.uz/oz/arkhiv-kursov-valyut/json/")

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	result := 0.00
	err = json.Unmarshal(body, &res)
	for _, val := range res {
		if val.Ccy == req.Ccy {
			intg, err := strconv.ParseFloat(val.Rate, 64)
			if err != nil {
				fmt.Println("error while converting rate string to integer")
			}
			mon, err := strconv.ParseFloat(req.Money, 64)

			if err != nil {
				fmt.Println("error while converting money string to integer")
			}
			result += (mon / intg)
		}
	}

	x := int(math.Ceil(result))
	final := (strconv.Itoa(x))

	return &dict.MoneyConverterResponse{
		Result: final,
		Ccy:    req.Ccy,
	}, nil

}
func (s *server) Square(ctx context.Context, req *dict.NumberRequest) (*dict.NumberResponse, error) {

	result := math.Pow(float64(req.Number), float64(req.Degree))
	// ress, err := strconv.ParseInt(result, 64)
	return &dict.NumberResponse{
		Result: int64(result),
	}, nil
}

func (s *server) MaxNum(ctx context.Context, req *dict.MaxRequest) (*dict.MaxResponse, error) {

	Largest := req.Nums[0]
	for _, val := range req.Nums {
		if val > Largest {
			Largest = val
		}
	}
	return &dict.MaxResponse{
		Result: Largest,
	}, nil

}
func main() {

	lis, err := net.Listen("tcp", "localhost:9101")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	dict.RegisterTranslateServer(s, &server{})

	fmt.Println("Listenining RPC SERVER...", "localhost:9101")
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
