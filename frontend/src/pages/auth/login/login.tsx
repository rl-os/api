import React, { useEffect } from 'react';
import classNames from "classnames";
import { observer } from 'mobx-react-lite';
import { Link } from 'react-router-dom';
import { useForm } from "react-hook-form";
import { useStore } from '../../../store';
import { useRouter } from "../../../utils/hooks";
import Noty from "noty";

import styles from './login.module.scss';
import { Logo } from './logo';

const Login = () => {
  const store = useStore();
  const router = useRouter();
  const { handleSubmit, register, errors } = useForm();

  useEffect(() => {
    if (store.auth.currentUser === null) { return; }
    router.history.push('/');
  }, [router.history, store.auth.currentUser])

  const onLogin = () => {
    new Noty({
      type: "success",
      text: `Вы успешно вошли в свой аккаунт!`,
      timeout: 5000,
    }).show();

    router.history.push('/');
  }
  const onError = (e: any) => {
    new Noty({
      type: "error",
      text: e.response?.data.message || e.message,
      timeout: 5000,
    }).show();
  };

  const onSubmit = ({username, password}: any) =>
    store.auth.login(username, password)
      .then(onLogin)
      .catch(onError);

  return <div className={classNames(styles.root)}>
    <div className={classNames(styles.aside, "d-flex flex-row-auto")}>
      <Logo/>
    </div>

    <div className="flex-row-fluid d-flex flex-column justify-content-center position-relative overflow-hidden p-5 ml-auto mr-auto">
      <div className="d-flex flex-column-fluid flex-center mt-6 mt-lg-0">
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className={classNames(styles.title)}>
            <h3 className="font-weight-bolder text-dark display5">
              Добро пожаловать на Risu.Life
            </h3>
            <span className="text-muted font-weight-bold font-size-h4">
              Нет аккаунта? <Link to="/auth/signup" className="text-primary font-weight-bolder">
                Регистрация
              </Link>
            </span>
          </div>

          <div className="form-group fv-plugins-icon-container">
            <label className="font-size-h6 font-weight-bolder text-dark">Логин</label>
            <input
              className="form-control form-control-solid h-auto py-7 px-6 rounded-lg"
              type="text"
              name="username"
              ref={register({
                required: "Необходимо заполнить",
                validate: value => value.length >= 3 ? true : 'Не верный логин'
              })}
            />
            {errors.username && errors.username.message}
          </div>

          <div className="form-group fv-plugins-icon-container">
            <label className="font-size-h6 font-weight-bolder text-dark">Пароль</label>
            <input
              className="form-control form-control-solid h-auto py-7 px-6 rounded-lg"
              type="password"
              name="password"
              ref={register({
                required: "Необходимо заполнить",
                validate: value => value.length >= 8 ? true : 'Пароль должен быть больше или равен 8 символам'
              })}
            />
            {errors.password && errors.password.message}
          </div>

          <button className={classNames(styles.btn, "btn btn-primary")} type="submit">
            Войти
          </button>

        </form>
      </div>

      <div className="d-flex justify-content-center align-items-end pb-5 pt-5">
        <Link to="/" className="text-primary font-weight-bolder font-size-h5">
          Правила
        </Link>
        <Link to="/" className="text-primary ml-3 font-weight-bolder font-size-h5">
          Контакты
        </Link>
        <Link to="/" className="text-primary ml-3 font-weight-bolder font-size-h5">
          О проекте
        </Link>
      </div>
    </div>
  </div>;
};

export default observer(Login);
