package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"

	"http_service/pkg/storage"
	"http_service/pkg/user"
)

type ApiUser struct {
	*storage.UserList
}

func NewApiUser(ctx context.Context, coll *mongo.Collection) *ApiUser {
	return &ApiUser{storage.NewUserList(ctx, coll)}
}
func (a *ApiUser) Run(port string) error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/", func(r chi.Router) {
		r.Get("/{userID}", a.GetOneUser)
		r.Post("/", a.CreateUser)
		r.Post("/make_friends", a.MakeFriends)
		r.Put("/{userID}", a.UpdateAge)
		r.Delete("/{userID}", a.DeleteUser)
		r.Route("/friends", func(r chi.Router) {
			r.Get("/{userID}", a.GetUserFriends)
		})
	})
	return http.ListenAndServe(":"+port, r)
}
func (a *ApiUser) GetUserFriends(w http.ResponseWriter, r *http.Request) {
	if userID := chi.URLParam(r, "userID"); userID != "" {
		userID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		slicefriend, err := a.GetUserFriend(userID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, a.FriensListResponse(slicefriend))
	}
}
func (a *ApiUser) GetOneUser(w http.ResponseWriter, r *http.Request) {
	if userID := chi.URLParam(r, "userID"); userID != "" {
		userID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		data, err := a.GetUser(userID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		render.Status(r, http.StatusOK)
		render.Render(w, r, data)
	}
}
func (a *ApiUser) FriensListResponse(u []*user.Friend) []render.Renderer {
	list := []render.Renderer{}
	
	for _, one_user := range u {
		list = append(list, one_user)
	}
	return list
}
func (a *ApiUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	newuser := &user.User{}
	if err := render.Bind(r, newuser); err != nil {
		fmt.Println(err)
		return
	}
	if err := a.AppendUser(newuser); err != nil {
		fmt.Println(err)
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, newuser)
}
func (a *ApiUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	data := &user.User{}
	if userID := chi.URLParam(r, "userID"); userID != "" {
		userID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		deluser, err := a.GetUser(userID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		data.Name = deluser.Name
		if err := a.DelUser(deluser.Id); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		render.Status(r, http.StatusOK)
		render.Render(w, r, data)
	}
}
func (a *ApiUser) MakeFriends(w http.ResponseWriter, r *http.Request) {
	data := &user.FriendRequest{}
	if err := render.Bind(r, data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	user_sourse, err := a.GetUser(data.Source_id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("несуществует пользователя с id %d", data.Source_id)))
		return
	}
	user_target, err := a.GetUser(data.Target_id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("несуществует пользователя с id %d", data.Target_id)))
		return

	}
	if err := user_sourse.NewFiend(user_target); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err := a.UpdateUser(user_sourse); err != nil {
		fmt.Println(err)
	}
	if err := a.UpdateUser(user_target); err != nil {
		fmt.Println(err)
	}
	render.Status(r, http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s и %s теперь друзья \n", user_sourse.Name, user_target.Name)))
}
func (a *ApiUser) UpdateAge(w http.ResponseWriter, r *http.Request) {
	newdata := &user.UserAgeUpdate{}
	if err := render.Bind(r, newdata); err != nil {
		fmt.Println(err)
		return
	}
	if userID := chi.URLParam(r, "userID"); userID != "" {
		userID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		updateuser, err := a.GetUser(userID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		updateuser.Age = newdata.NewAge
		if err := a.UpdateUser(updateuser); err != nil {
			fmt.Println(err)
		}
		w.Write([]byte("возраст пользователя успешно обнавлен"))
	}
}
