import React from 'react';
import { observer } from 'mobx-react-lite';
import { Link } from "react-router-dom";
import { useStore } from "../../store";
import { useTranslate } from "../../utils/hooks";

const Login = () => {
  const store = useStore();
  const { t } = useTranslate("auth");

  const onLogout = () =>
    store.auth.reset();

  return <div>
    {
      store.auth.currentUser
        ? <a href="#" onClick={onLogout}>{t`Log out`}</a>
        : <Link to="/auth/login">{t`Log in`}</Link>
    }
  </div>;
};

export default observer(Login);
