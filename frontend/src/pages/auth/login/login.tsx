import React from 'react';
import classNames from "classnames";
import { observer } from 'mobx-react-lite';

import styles from './login.module.scss';

const Login = () => {
  return <div className={classNames(styles.root)}>
    <div className={classNames(styles.aside, "d-flex flex-row-auto")}>
      text
    </div>

    <div className="flex-row-fluid d-flex flex-column justify-content-center position-relative overflow-hidden p-7 ml-auto mr-auto">
      <div className="d-flex flex-column-fluid flex-center mt-6 mt-lg-0">
        <form>
          <div className="pt-lg-40 mt-lg-10 pb-15">
            <h3 className="font-weight-bolder text-dark display5">Добро пожаловать на Risu.Life</h3>
            <span className="text-muted font-weight-bold font-size-h4">
            Нет аккаунта? <a href="#" className="text-primary font-weight-bolder">Регистрация</a>
          </span>
          </div>
        </form>
      </div>

      <div className="d-flex justify-content-lg-start justify-content-center flex-column-fluid align-items-end pb-2 pt-lg-0">
        <a href="#" className="text-primary font-weight-bolder font-size-h5">Правила</a>
        <a href="#" className="text-primary ml-3 font-weight-bolder font-size-h5">Контакты</a>
        <a href="#" className="text-primary ml-3 font-weight-bolder font-size-h5">О проекте</a>
      </div>
    </div>
  </div>;
};

export default observer(Login);
