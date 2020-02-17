import history from 'history';
import React, { createElement, useEffect, useState } from 'react';
import { Redirect, Route, RouteComponentProps, RouteProps } from 'react-router-dom';

interface BasicState {
  status: 'loading' | 'success';
}

interface RedirectState {
  status: 'redirect';
  url: string;
}

type RouteState = BasicState | RedirectState;

function checkGuard(location: history.Location, props: SecureRouteProps, setState: (state: RouteState) => void) {
  let cancelled = false;
  (async () => {
    let state: RouteState = { status: 'loading' };
    try {
        const redirect = props.guard && await props.guard(location);
        state = redirect !== undefined ? { status: 'redirect', url: redirect } : { status: 'success' };
      } catch (e) {
        console.error('Error in guard', e);
      }
    if (!cancelled) {
        setState(state);
      }
  })();
  return () => {
    cancelled = true;
  };
}

interface WrapperProps {
  inner: RouteComponentProps<any>;
  outer: SecureRouteProps;
}

function RouteWrapper({ inner, outer }: WrapperProps) {
  const [state, setState] = useState<RouteState>({
    status: 'loading',
  });
  useEffect(() => checkGuard(inner.location, outer, setState), [inner.location, outer]);

  if (state.status === 'redirect') {
    return <Redirect to={state.url}/>;
  }

  if (state.status === 'success') {
    if (outer.component) {
        return createElement(outer.component, inner);
      }
  }

  return <div/>;
}

export type GuardResult = string | undefined;

export interface SecureRouteProps extends RouteProps {
  guard?: (location: history.Location) => Promise<GuardResult> | GuardResult;
}

export const SecureRoute = (props: SecureRouteProps) => {
  const render = (inner: RouteComponentProps<any>) => <RouteWrapper inner={inner} outer={props}/>;
  return <Route path={props.path} exact={props.exact} render={render}/>;
};
