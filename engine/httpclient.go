package engine

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type eventClient struct {
}

type resPack struct {
	r   *http.Response
	err error
}

func (ec *eventClient) work(ctx context.Context, req *http.Request) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan resPack, 1)
	go func() {
		resp, err := client.Do(req)
		pack := resPack{r: resp, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		fmt.Println("Timeout!")
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}
