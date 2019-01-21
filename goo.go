//goo.go this file incloud the common functions,maps slices, structs or anything else.
package goo

import (
	"cloud.google.com/go/translate"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

// InSet is a bit array struct.
type IntSet struct {
	Words []uint64 // bit array
}

//Len return a int type value about the bit array's length.
func (i *IntSet) Len() int {
	return len(i.Words)
}

// remove this array's one bit.
func (i *IntSet) Remove(x int) {
	if i.Len() <= x {
		os.Exit(1)
	}
	i.Words = append(i.Words[:x], i.Words[x+1:]...)
}

// clear the array,and let  i.Words point a new slice.
func (i *IntSet) Clear() {
	i.Words = make([]uint64, 0)
}

// copy form the old array to a new array,return a new instance's pointer.
func (i *IntSet) Copy() *IntSet {
	t := *i
	return &t
}

// ip

type Data struct {
	Data Values `json:"data"`
}
type Values struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

// add ip address, return country and city.
func WhichCountry(ip ...string) (values []*Values) {
	var syGroup sync.WaitGroup
	var sySync sync.Mutex
	vi := new(Data)
	syGroup.Add(len(ip))
	controlSpeed := make(chan struct{}, 20) // control http-get's speed.
	for _, v := range ip {
		go func(ip string) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()
			defer syGroup.Done()
			controlSpeed <- struct{}{} // 启动计数器
			res, _ := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
			<-controlSpeed // 结束计数器
			defer res.Body.Close()
			data, _ := ioutil.ReadAll(res.Body)
			json.Unmarshal(data, vi)
			sySync.Lock()
			values = append(values, &(vi.Data))
			sySync.Unlock()
		}(v)
	}
	syGroup.Wait()
	return
}

func WhichCountryInLanguages(language string, ip ...string) (values []*Values) {
	value := WhichCountry(ip...)
	vi := new(Values)
	for _, v := range value {
		Country, err := setLanguage(language, v.Country)
		if err != nil {
			fmt.Println(err)
		}
		City, err := setLanguage(language, v.City)
		if err != nil {
			fmt.Println(err)
		}
		vi.Country = Country
		vi.City = City
		values = append(values, vi)
	}
	return
}

func setLanguage(lan, text string) (value string, err error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}

	target, err := language.Parse(lan)
	if err != nil {
		return "", err
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Translation: %v\n", translations[0].Text), nil
}
