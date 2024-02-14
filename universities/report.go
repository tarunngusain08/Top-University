package universities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type report struct {
	mu              *sync.Mutex
	wg              *sync.WaitGroup
	pageNum         int
	responseChannel chan []*University
	numWorkers      int
	done            bool
}

type Universities interface {
	HighestInternationalStudents(firstCity string, secondCity string) (string, error)
}

func NewReport() Universities {
	return &report{
		mu:              new(sync.Mutex),
		wg:              new(sync.WaitGroup),
		pageNum:         0,
		responseChannel: make(chan []*University),
		numWorkers:      10,
	}
}

type University struct {
	UniversityName           string    `json:"UniversityName"`
	RankDisplay              string    `json:"rank_display"`
	Score                    float64   `json:"score"`
	Type                     string    `json:"type"`
	StudentFacultyRatio      int       `json:"student_faculty_ratio"`
	NumInternationalStudents string    `json:"international_students"`
	FacultyCount             string    `json:"faculty_count"`
	Location                 *location `json:"location"`
}

type location struct {
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

func (r *report) producer() {
	defer r.wg.Done()
	if r.done {
		return
	}
	r.mu.Lock()
	r.pageNum++
	r.mu.Unlock()
	fmt.Printf("Fetching %d page..\n", r.pageNum)
	res, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/universities?page=%v", r.pageNum))
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()
	var resp UniversityData
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("producer sending...", resp.Data)
	r.responseChannel <- resp.Data
}

func (r *report) consumer(AllUniversities *[]*University) {
	for {
		select {
		case universities, ok := <-r.responseChannel:
			if !ok {
				break
			}
			fmt.Println("consumer reading from channel...", universities)
			*AllUniversities = append(*AllUniversities, universities...)
		}
	}
}

func (r *report) HighestInternationalStudents(firstCity string, secondCity string) (string, error) {
	defer close(r.responseChannel)

	for i := 0; i < 20; i++ {
		r.wg.Add(1)
		go r.producer()
	}
	universities := make([]*University, 0)
	go r.consumer(&universities)

	r.wg.Wait()

	fmt.Println(universities)
	winner := new(University)
	winner.NumInternationalStudents = "0"

	for _, university := range universities {
		if university.Location.City == firstCity || university.Location.City == secondCity {
			winnerNum, err := strconv.Atoi(r.convertString(winner.NumInternationalStudents))
			if err != nil {
				return "", err
			}
			universityNum, err := strconv.Atoi(r.convertString(university.NumInternationalStudents))
			if err != nil {
				return "", err
			}
			if winnerNum < universityNum {
				winner = &University{
					UniversityName:           university.UniversityName,
					RankDisplay:              university.RankDisplay,
					Score:                    university.Score,
					Type:                     university.Type,
					StudentFacultyRatio:      university.StudentFacultyRatio,
					NumInternationalStudents: university.NumInternationalStudents,
					FacultyCount:             university.FacultyCount,
					Location:                 university.Location,
				}
			}
		}
	}
	return winner.UniversityName, nil
}

func (r *report) convertString(s string) string {
	ans := ""
	for _, i := range s {
		if i >= '0' && i <= '9' {
			ans += string(i)
		}
	}
	return ans
}
