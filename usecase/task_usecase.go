package usecase

import (
	"github.com/todo_test/model"
	"github.com/todo_test/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskRespons, error)
	GetTaskById(userId uint, taskId uint) (model.TaskRespons, error)
	CreateTask(task model.Task) (model.TaskRespons, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskRespons, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskRespons, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskRespons{}
	for _, v := range tasks {
		t := model.TaskRespons{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskRespons, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskRespons{}, err
	}
	resTask := model.TaskRespons{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskRespons, error) {
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskRespons{}, nil
	}
	resTask := model.TaskRespons{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil

}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskRespons, error) {
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskRespons{}, err
	}
	resTask := model.TaskRespons{
		ID:        taskId,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
