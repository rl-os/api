import React from 'react';
import { observer } from 'mobx-react-lite';
import { Link } from "react-router-dom";

const Login = () => {
  return <div>
    <Link to="/login">Login</Link>
  </div>;
};

export default observer(Login);
