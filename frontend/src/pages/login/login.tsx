import React from 'react';
import classNames from 'classnames';
import { observer } from 'mobx-react-lite';

import styles from './login.module.scss';

const Login = () => {
  return <div className={classNames(styles.root)}>
    <nav className="navbar navbar-expand-lg navbar-light bg-dark fixed-top">
      <div className="container">
        <img src="https://301222.selcdn.ru/akasi/assets/logo/26.04-2020.psd.jpg" width="40" height="40" alt=""/>
      </div>
    </nav>

    <div className="container">
      <div className={classNames(styles.container, "row justify-content-center align-items-center")}>
        <div className="col-12 col-md-5">
          <div className="card login">
            <div className="card-body">
              <div className="form-group">
                <label htmlFor="password">Логин</label>
                <input type="email" className="form-control" required/>
              </div>

              <div className="form-group">
                <label htmlFor="password2">Пароль</label>
                <input type="password" className="form-control" required/>
              </div>

              <button className="btn btn-primary btn-sm submit-btn" id="change" value="Войти">
                Сменить
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>;
};

export default observer(Login);
