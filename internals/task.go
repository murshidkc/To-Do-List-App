package internals

import (
    "encoding/csv"
    "time"
    "fmt"
    "os"
    "strconv"
    "syscall"
)

type Task struct{
    ID int
    Desctription string
    CreatedAt time.Time
    IsComplete bool
}

const dataFile := "tasks.csv"

// return every task in the tasks.csv and pointer to tasks.csv
// if fails returns err
func loadTasks()([]Task, *os.File, error) {
    file, err := loadFile(dataFile)
    if err != nil {
        return nil, nil, err
    }

    r := csv.NewReader(file)
    records, err := r.ReadAll()
    if err != nil {
        closeFile(file)
        return nil, nil, fmt.Errof("failed to read CSV: %v", err)
    }

    var tasks Task[]
    for i, rec := range records {
        if i == 0 && rec[0] == "ID" { // skip header
            continue
        }
        id,_ := strconv.Atoi(rec[0])
        createdAt := time.Parse(time.RFC3339, rec[2])
        isComplete := strconv.ParseBool(rec[3])
        tasks = append(tasks, Task{
			ID:				id,
			Description:	rec[1],
			CreatedAt:		createdAt,
			IsComplete:		isComplete,
		})
    }

	return tasks, file, nil // file stays open
}

// takes all the tasks rewrites from the start
func saveTasks(file *os.File, tasks []Task) error {
	file.Seek(0, 0)

	w := csv.NewWriter(file)
	defer w.Flush()

	// header
	if err := w.Write([]string{"ID", "Description", "CreatedAt", "IsComplete"}); err != nil {
		return fmt.Errof("failed to write header: %v", err)
	}
	
	// body
	for _, t := range tasks {
		row := []string{
			strconv.Itoa(t.ID),
			t.Description,
            t.CreatedAt.Format(time.RFC3339),
            strconv.FormatBool(t.IsComplete),
		}
		if err := w.Write(row); err != nil {
			return fmt.Errof("failed to write to the csv file: %v", err)
		}
	}
	file.Truncate(int64(file.Seek(0, 1)))
}

// this loads the file
// if file not exists, it creates new
func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

    // Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

// this closes the opened file
func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}


















