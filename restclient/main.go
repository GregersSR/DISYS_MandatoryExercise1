package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	baseurl := "http://localhost:8080/course"

	fmt.Println("Starting REST client")

	// Post a couple of courses
	postCourses(baseurl)

	// Let's see how the course with ID 2 looks
	fmt.Println("Course 2:")
	body, err := getBody(baseurl + "/2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body)
	}

	// Now let's change some stuff; we add a teacher and raise the satisfaction score
	fmt.Println("Changing course with ID 2")
	teachers := make([]int64, 0)
	teachers = append(teachers, 1)
	resp2, err := doJSONRequest(baseurl+"/2", "PUT", map[string]interface{}{
		"id":                 2,
		"title":              "Discrete Mathematics",
		"ects":               10,
		"satisfaction_score": 95,
		"teachers":           teachers,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp2.Status)
		resp2.Body.Close()
	}

	// How does it look now?
	fmt.Println("Course 2:")
	body2, err := getBody(baseurl + "/2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body2)
	}

	fmt.Println("Teachers of course 2:")
	body3, err := getBody(baseurl + "/2/teachers")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body3)
	}

	fmt.Println("Satisfaction score of course 2:")
	body4, err := getBody(baseurl + "/2/satisfaction_rating")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body4)
	}

	// Let's clean up
	fmt.Println("Deleting course with ID 1")
	resp4, err := doJSONRequest(baseurl+"/1", "DELETE", map[string]interface{}{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp4.Status)
	}
	fmt.Println("Deleting course with ID 2")
	resp5, err := doJSONRequest(baseurl+"/2", "DELETE", map[string]interface{}{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp5.Status)
	}
	fmt.Println("Attepting delete of non-existing course with ID 3")
	resp6, err := doJSONRequest(baseurl+"/3", "DELETE", map[string]interface{}{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp6.Status)
	}

}

func postCourses(baseurl string) {
	fmt.Println("Creating course Distributed Systems with ID 1")
	resp, err := doJSONRequest(baseurl, "POST", map[string]interface{}{
		"id":                 1,
		"title":              "Distributed Systems",
		"ects":               7,
		"satisfaction_score": 90,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Status)
	}

	fmt.Println("Creating course Discrete Mathematics with ID 2")
	resp, err = doJSONRequest(baseurl, "POST", map[string]interface{}{
		"id":                 2,
		"title":              "Discrete Mathematics",
		"ects":               10,
		"satisfaction_score": 80,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Status)
	}
}

func getBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), nil
}

func doJSONRequest(url string, method string, dataToSend map[string]interface{}) (*http.Response, error) {
	asJSON, err := json.Marshal(dataToSend)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(asJSON))
	req.Header.Set("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
