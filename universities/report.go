package universities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Report struct {
	mu              sync.Mutex
	responseChannel chan []*University
}

type University struct {
	UniversityName           string    `json:"UniversityName"`
	RankDisplay              string    `json:"rank_display"`
	Score                    float64   `json:"score"`
	Type                     string    `json:"type"`
	StudentFacultyRatio      int       `json:"student_faculty_ratio"`
	NumInternationalStudents string    `json:"international_students"`
	FacultyCount             string    `json:"faculty_count"`
	Location                 *Location `json:"location"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Region  string `json:"region"`
}

type UniversityData struct {
	Page       int           `json:"page"`
	PerPage    int           `json:"per_page"`
	Total      int           `json:"total"`
	TotalPages int           `json:"total_pages"`
	Data       []*University `json:"data"`
}

func NewReport() *Report {
	return &Report{
		responseChannel: make(chan []*University),
	}
}

func (r *Report) FetchUniversityData(pageNum int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Fetching page %d...\n", pageNum)
	resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/universities?page=%d", pageNum))
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var universityData UniversityData
	if err := json.NewDecoder(resp.Body).Decode(&universityData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	r.responseChannel <- universityData.Data
}

func (r *Report) HighestInternationalStudents(firstCity, secondCity string) (string, error) {
	defer close(r.responseChannel)

	const numPages = 20
	var wg sync.WaitGroup
	wg.Add(numPages)

	for i := 1; i <= numPages; i++ {
		go r.FetchUniversityData(i, &wg)
	}

	wg.Wait()

	var winner *University
	maxStudents := 0

	for universities := range r.responseChannel {
		for _, u := range universities {
			if u.Location.City == firstCity || u.Location.City == secondCity {
				numStudents, err := strconv.Atoi(strings.ReplaceAll(u.NumInternationalStudents, ",", ""))
				if err != nil {
					return "", err
				}
				if numStudents > maxStudents {
					maxStudents = numStudents
					winner = u
				}
			}
		}
	}

	if winner == nil {
		return "", fmt.Errorf("no universities found in the specified cities")
	}

	return winner.UniversityName, nil
}
