package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"my-reading-app/internal/domain"
	"os"
	"strconv"
	"strings"
)

var bookMap = map[string]string{
	"Genesis":         "gn",
	"Exodus":          "ex",
	"Leviticus":       "lv",
	"Numbers":         "nm",
	"Deuteronomy":     "dt",
	"Joshua":          "js",
	"Judges":          "jz",
	"Ruth":            "rt",
	"1 Samuel":        "1sm",
	"2 Samuel":        "2sm",
	"1 Kings":         "1rs",
	"2 Kings":         "2rs",
	"1 Chronicles":    "1cr",
	"2 Chronicles":    "2cr",
	"Ezra":            "esd",
	"Nehemiah":        "ne",
	"Tobit":           "tb",
	"Judith":          "jt",
	"Esther":          "est",
	"1 Maccabees":     "1mc",
	"2 Maccabees":     "2mc",
	"Job":             "jó",
	"Psalm":           "sl",
	"Proverbs":        "pr",
	"Ecclesiastes":    "ecl",
	"Song of Songs":   "ct",
	"Wisdom":          "sb",
	"Sirach":          "eclo",
	"Isaiah":          "is",
	"Jeremiah":        "jr",
	"Lamentations":    "lm",
	"Baruch":          "br",
	"Ezekiel":         "ez",
	"Daniel":          "dn",
	"Hosea":           "os",
	"Joel":            "jl",
	"Amos":            "am",
	"Obadiah":         "ob",
	"Jonah":           "jn",
	"Micah":           "mq",
	"Nahum":           "na",
	"Habakkuk":        "hc",
	"Zephaniah":       "sf",
	"Haggai":          "ag",
	"Zechariah":       "zc",
	"Malachi":         "ml",
	"Matthew":         "mt",
	"Mark":            "mc",
	"Luke":            "lc",
	"John":            "jo",
	"Acts":            "at",
	"Romans":          "rm",
	"1 Corinthians":   "1cor",
	"2 Corinthians":   "2cor",
	"Galatians":       "gl",
	"Ephesians":       "ef",
	"Philippians":     "fp",
	"Colossians":      "cl",
	"1 Thessalonians": "1ts",
	"2 Thessalonians": "2ts",
	"1 Timothy":       "1tm",
	"2 Timothy":       "2tm",
	"Titus":           "tt",
	"Philemon":        "fm",
	"Hebrews":         "hb",
	"James":           "tg",
	"1 Peter":         "1pd",
	"2 Peter":         "2pd",
	"1 John":          "1jo",
	"2 John":          "2jo",
	"3 John":          "3jo",
	"Jude":            "jd",
	"Revelation":      "ap",
}

type BibleService interface {
	GetBibleText(description string) ([]string, error)
}

type bibleService struct {
	bibleData []domain.BibleVerse
}

func NewBibleService(filePath string) (BibleService, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return nil, err
	}

	var bibleData []domain.BibleVerse
	err = json.Unmarshal(data, &bibleData)
	if err != nil {
		log.Printf("Error unmarshalling data: %v", err)
		return nil, err
	}

	log.Println("Bible data loaded successfully")
	return &bibleService{bibleData: bibleData}, nil
}

func getShortName(bookName string) (string, error) {
	// Check for multi-word book names like "1 Samuel"
	if len(strings.Split(bookName, " ")) > 1 {
		for fullName, shortName := range bookMap {
			if strings.HasPrefix(strings.ToLower(fullName), strings.ToLower(bookName)) {
				return shortName, nil
			}
		}
	}
	// Check for single-word book names
	for fullName, shortName := range bookMap {
		if strings.EqualFold(fullName, bookName) {
			return shortName, nil
		}
	}
	return "", errors.New("book not found")
}

func (b *bibleService) GetBibleText(description string) ([]string, error) {
	log.Printf("Fetching bible text for description: %s", description)
	parts := strings.Split(description, " ")
	if len(parts) < 2 {
		log.Println("Invalid description format")
		return nil, errors.New("invalid description format")
	}

	bookName := parts[0]
	if len(parts) > 2 {
		bookName = parts[0] + " " + parts[1]
		parts = append(parts[:1], parts[2:]...)
	}
	chaptersOrVerses := parts[1]

	bookShortName, err := getShortName(bookName)
	if err != nil {
		log.Println("Book not found")
		return nil, err
	}

	var result []string
	if strings.Contains(chaptersOrVerses, ":") {
		// Handle verse range like "Proverbs 2:9-15"
		chapterVerseParts := strings.Split(chaptersOrVerses, ":")
		if len(chapterVerseParts) != 2 {
			log.Println("Invalid verse range format")
			return nil, errors.New("invalid verse range format")
		}

		chapter, err := strconv.Atoi(chapterVerseParts[0])
		if err != nil {
			log.Println("Invalid chapter")
			return nil, errors.New("invalid chapter")
		}

		verseRange := chapterVerseParts[1]
		verseParts := strings.Split(verseRange, "-")
		if len(verseParts) != 2 {
			log.Println("Invalid verse range")
			return nil, errors.New("invalid verse range")
		}

		startVerse, err := strconv.Atoi(verseParts[0])
		if err != nil {
			log.Println("Invalid start verse")
			return nil, errors.New("invalid start verse")
		}
		endVerse, err := strconv.Atoi(verseParts[1])
		if err != nil {
			log.Println("Invalid end verse")
			return nil, errors.New("invalid end verse")
		}

		for _, verse := range b.bibleData {
			if strings.EqualFold(verse.Livro, bookShortName) && verse.Capitulo == chapter {
				for i := startVerse; i <= endVerse; i++ {
					result = append(result, verse.Versiculos[i-1])
				}
				log.Printf("Fetched bible text: %v", result)
				return result, nil
			}
		}
	} else {
		// Handle chapter range like "Genesis 22-23"
		chapterParts := strings.Split(chaptersOrVerses, "-")
		if len(chapterParts) == 1 {
			// Single chapter
			chapter, err := strconv.Atoi(chapterParts[0])
			if err != nil {
				log.Println("Invalid chapter")
				return nil, errors.New("invalid chapter")
			}
			for _, verse := range b.bibleData {
				if strings.EqualFold(verse.Livro, bookShortName) && verse.Capitulo == chapter {
					log.Printf("Fetched bible text: %v", verse.Versiculos)
					return verse.Versiculos, nil
				}
			}
		} else {
			// Multiple chapters
			startChapter, err := strconv.Atoi(chapterParts[0])
			if err != nil {
				log.Println("Invalid start chapter")
				return nil, errors.New("invalid start chapter")
			}
			endChapter, err := strconv.Atoi(chapterParts[1])
			if err != nil {
				log.Println("Invalid end chapter")
				return nil, errors.New("invalid end chapter")
			}
			for _, verse := range b.bibleData {
				if strings.EqualFold(verse.Livro, bookShortName) && verse.Capitulo >= startChapter && verse.Capitulo <= endChapter {
					result = append(result, verse.Versiculos...)
				}
			}
			log.Printf("Fetched bible text: %v", result)
			return result, nil
		}
	}
	log.Println("Text not found")
	return nil, errors.New("text not found")
}
