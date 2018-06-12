package test
import (
	"fmt"
	"os"
	"testing"
	"tmdb/parser"
	"io/ioutil"
	"log"
)
func TestHeaderParsing(t *testing.T) {
	file, _ := os.Open("tmdb_5_credits.csv")
	movies := parser.ParseCsvFile(file)
	nr_movies := 4
	if (len(movies) != nr_movies) {
		t.Errorf("should be nr of movies %d but is %d", nr_movies, len(movies))
	}
}
func createCsv(nrOfRecords int) (*os.File, error) {
	header := "movie_id,title,cast,crew\n"
	actor := "[{\"\"cast_id\"\":%d, \"\"character\"\":\"\"mycharacter\"\", \"\"name\"\":\"\"somename\"\"}]"
	movie := "%d,Title of movie %d,\"%s\",\"[]\"\n"

	tmpfile, err := ioutil.TempFile(os.TempDir(), "test_movie.csv")
	if _, err := tmpfile.WriteString(header); err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()
	for index := 0; index < nrOfRecords; index++ {
		formattedActor := fmt.Sprintf(actor, index)
		formattedMovie := fmt.Sprintf(movie, index, index, formattedActor)
		_, err := tmpfile.WriteString(formattedMovie)
		if (err != nil) {
			log.Fatal(err)
		}
	}
	
	return tmpfile, err
}
func TestMovie(t *testing.T) {
	
	tmpfile, _ := createCsv(3)
	csvFile, _ := os.Open(tmpfile.Name())

	defer os.Remove(tmpfile.Name()) // clean up
	
	movies := parser.ParseCsvFile(csvFile)
	nr_movies := 3
	if (len(movies) != nr_movies) {
		t.Errorf("should be nr of movies %d but is %d", nr_movies, len(movies))
	}
}

func BenchmarkMovie(b *testing.B) {
	
	tmpfile, _ := createCsv(b.N)
	csvFile, _ := os.Open(tmpfile.Name())

	defer os.Remove(tmpfile.Name()) // clean up
	
	movies := parser.ParseCsvFile(csvFile)
	nr_movies := b.N
	if (len(movies) != nr_movies) {
		b.Errorf("should be nr of movies %d but is %d", nr_movies, len(movies))
	}
}