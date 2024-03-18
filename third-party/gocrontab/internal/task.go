package internal

type task struct {

	//the command will be runing
	cmd string
	//the task of command
	comment string
	//the minute will be runing 0-59
	minutes map[int]struct{}
	//the hours will be runing 0-23
	hours map[int]struct{}
	//the months will be runing 1-12
	months map[int]struct{}
	//the weekday of crontab
	weekDay map[int]struct{}
}

func (t *task) Run() error {

	return nil
}

func (t *task) Stop() error {

	return nil
}
