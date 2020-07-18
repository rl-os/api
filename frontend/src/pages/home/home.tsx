import React from 'react';
import { observer } from 'mobx-react-lite';
import { Link } from "react-router-dom";
import { useStore } from "../../store";
import { useTranslation } from "react-i18next";

const Login = () => {
  const store = useStore();
  const { t } = useTranslation();

  const onLogout = () =>
    store.auth.reset();

  return <div>
    {
      store.auth.currentUser
        ? <a href="#" onClick={onLogout}>{t`auth:Log in`}</a>
        : <Link to="/auth/login">Login</Link>
    }
  </div>;
};

export default observer(Login);
