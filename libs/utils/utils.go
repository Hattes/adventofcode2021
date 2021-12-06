package utils;

import (
  "fmt"
  "os"
  "errors"
  "strings"
  "strconv"
  "net/http"
  "io/ioutil"
  "path/filepath"
);

const url_format string = "https://adventofcode.com/%d/day/%d/input";
const path_session_token = "cookie_session";

func get_session() (string, error) {
    content, err := ioutil.ReadFile(path_session_token)
    if err != nil {
        fmt.Fprintln(os.Stderr, "ERROR (utils.get_session): Failed to read ./"+path_session_token);
        return "", err
    }

    return strings.TrimSpace(string(content)), nil;
}

func Get_input(year int, day int) (string, error) {
    local_data, err := get_input_from_local_cache(day)
    if err == nil {
        return string(local_data), nil
    }
    return get_input_from_aoc(year, day)
}

func get_input_from_local_cache(day int) ([]byte, error) {
    file_name := fmt.Sprintf("input_%d.txt", day)
    project_root := os.Getenv("PROJECT_ROOT")  // Set in calling script
    path := filepath.Join(project_root, "inputs", file_name)
    return os.ReadFile(path)
}

func get_input_from_aoc(year int, day int) (string, error) {
    var session, err = get_session();

    // Sanity and errors
    if err != nil {
        return "", err;
    }
    if session == "" {
        const msg string = "ERROR (utils.Get_input): Session-token is an empty string. Make sure you edited ./cookie_session";
        fmt.Fprintln(os.Stderr, msg);
        return "", errors.New(msg);
    }

    var url = fmt.Sprintf(url_format, year, day);

    // Declare http client
    var client = &http.Client{};

    // Declare HTTP Method and Url
    req, err := http.NewRequest("GET", url, nil);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    // Set cookie
    req.Header.Set("Cookie", fmt.Sprintf("session=%s; count=x", session));

    resp, err := client.Do(req);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }
    // Read response
    data, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }
    cache_input(data, day)

    return string(data), nil;
}

func cache_input(input []byte, day int) error {
    file_name := fmt.Sprintf("input_%d.txt", day)
    project_root := os.Getenv("PROJECT_ROOT")  // Set in calling script
    path := filepath.Join(project_root, "inputs", file_name)
    return write_file(path, input)
}

func write_file(path string, input []byte) error {
    os.MkdirAll(filepath.Dir(path), os.ModePerm)
    err := os.WriteFile(path, input, 0644)
    if err != nil {
        fmt.Printf("error = %s \n", err);
        return err
    }
    return nil
}

func Trim_array(strs []string) []string {
    for i, str := range strs {
        strs[i] = strings.Trim(str, " ");
    }
    return strs;
}

func StrToInt_array(strs []string) ([]int, error) {
    var ints = make([]int, len(strs));
    for i, str := range strs {
        var d, err = strconv.Atoi(strings.TrimSpace(str));
        if err != nil {
            return nil, err;
        }
        ints[i] = d;
    }
    return ints, nil;
}

func StrToFloat_array(strs []string) ([]float64, error) {
    var floats = make([]float64,len(strs));
    for i, str := range strs {
        var f, err = strconv.ParseFloat(str, 64);
        if err != nil {
            return nil, err;
        }
        floats[i] = f;
    }
    return floats, nil;
}
func IntPow(n, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
