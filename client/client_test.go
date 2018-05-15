package client

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/sfluor/blazedb/server"
)

var testConfig = &server.Config{
	Port:         0,
	MaxQueueSize: 100,
	SaveFile:     "/tmp/blazedb.test.dump",
	LogFile:      "/tmp/blazedb.test.log",
	Debug:        0,
	SavePeriod:   999 * time.Hour,
}

func wrapper(f func(*Client)) {
	srv := server.New(testConfig)
	client := New(fmt.Sprintf("localhost:%v", srv.GetPort()))
	go srv.Start()
	for !srv.UP {
	}

	f(client)

	srv.Stop()
}

func TestValidSetThenGet(t *testing.T) {
	wrapper(func(client *Client) {

		tt := map[string][]byte{
			"1": []byte("{\"age\": 43, \"name\": \"john\"}"),
			"2": []byte("hey"),
			"3": []byte("{\"age\": 43, \"name\": \"john\"}"),
		}

		for key, value := range tt {
			err := client.Set(key, value)

			if err != nil {
				t.Errorf("Set command failed for (%s, %s): %s", key, value, err)
			}

			data, err := client.Get(key)

			if err != nil {
				t.Errorf("Get command failed for (%s, %s): %s", key, value, err)
			}

			if !reflect.DeepEqual(data, value) {
				t.Errorf("Get failed, expected %s, but got %s", value, data)
			}
		}
	})
}

func TestValidSetThenUpdate(t *testing.T) {
	wrapper(func(client *Client) {

		tt := []struct {
			key string
			old []byte
			new []byte
		}{
			{"1", []byte("{\"age\": 43, \"name\": \"john\"}"), []byte("hey")},
			{"2", []byte("hey"), []byte("hey")},
			{"3", []byte("{\"age\": 43, \"name\": \"john\"}"), []byte("hello")},
		}

		for _, tc := range tt {
			err := client.Set(tc.key, tc.old)

			if err != nil {
				t.Errorf("Set command failed for (%s, %s): %s", tc.key, tc.old, err)
			}

			err = client.Update(tc.key, tc.new)

			if err != nil {
				t.Errorf("Update command failed for (%s, %s): %s", tc.key, tc.new, err)
			}

			data, err := client.Get(tc.key)

			if !reflect.DeepEqual(data, tc.new) {
				t.Errorf("Get failed, expected %s, but got %s", tc.new, data)
			}
		}
	})
}

func TestValidSetThenDelete(t *testing.T) {
	wrapper(func(client *Client) {

		tt := map[string][]byte{
			"1": []byte("{\"age\": 43, \"name\": \"john\"}"),
			"2": []byte("hey"),
			"3": []byte("{\"age\": 43, \"name\": \"john\"}"),
		}

		for key, value := range tt {
			err := client.Set(key, value)

			if err != nil {
				t.Errorf("Set command failed for (%s, %s): %s", key, value, err)
			}

			err = client.Delete(key)

			if err != nil {
				t.Errorf("Delete command failed for (%s): %s", key, err)
			}

			data, err := client.Get(key)

			if err == nil {
				t.Errorf("Expected to encounter an error for get on %s but got '%s'", key, data)
			}

		}
	})
}

func TestInvalidGet(t *testing.T) {
	wrapper(func(client *Client) {

		tt := []string{
			"1", "2", "3",
		}

		for _, key := range tt {
			data, err := client.Get(key)

			if err == nil {
				t.Errorf("Expected to encounter an error for get on %s but got '%s'", key, data)
			}

		}
	})
}

func TestInvalidUpdate(t *testing.T) {
	wrapper(func(client *Client) {

		tt := map[string][]byte{
			"1": []byte("{\"age\": 43, \"name\": \"john\"}"),
			"2": []byte("hey"),
			"3": []byte("{\"age\": 43, \"name\": \"john\"}"),
		}

		for key, value := range tt {
			err := client.Update(key, value)

			if err == nil {
				t.Errorf("Expected to encounter an error for get on %s", key)
			}

		}
	})
}

func TestInvalidDelete(t *testing.T) {
	wrapper(func(client *Client) {

		tt := []string{
			"1", "2", "3",
		}

		for _, key := range tt {
			err := client.Delete(key)

			if err == nil {
				t.Errorf("Expected to encounter an error for get on %s", key)
			}

		}
	})
}

func TestInvalidSet(t *testing.T) {
	wrapper(func(client *Client) {

		tt := []struct {
			key string
			old []byte
			new []byte
		}{
			{"1", []byte("{\"age\": 43, \"name\": \"john\"}"), []byte("hey")},
			{"2", []byte("hey"), []byte("hey")},
			{"3", []byte("{\"age\": 43, \"name\": \"john\"}"), []byte("hello")},
		}

		for _, tc := range tt {
			err := client.Set(tc.key, tc.old)

			if err != nil {
				t.Errorf("Set command failed for (%s, %s): %s", tc.key, tc.old, err)
			}

			err = client.Set(tc.key, tc.new)

			if err == nil {
				t.Errorf("Expected to encounter an error for set on %s", tc.key)
			}
		}
	})
}
