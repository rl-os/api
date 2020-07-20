/* eslint-disable */
import { DependencyList, MutableRefObject, ReactElement, useCallback, useContext, useEffect, useRef } from 'react';
import { useAsync as useAsyncInner } from 'react-async';
import { __RouterContext } from 'react-router';
import {useTranslation, UseTranslationResponse} from 'react-i18next';

export function useRouter() {
  return useContext(__RouterContext);
}

export interface UseAsyncOptions<R> {
  init: () => Promise<R>;
  loading: () => ReactElement;
  success: (result: R) => ReactElement;
  failed: (error: any) => ReactElement;
  dependencies?: DependencyList;
}

export function useAsync<R>(options: UseAsyncOptions<R>): ReactElement {
  const init = useCallback(options.init, options.dependencies || []);

  const { data, error, isLoading } = useAsyncInner({
    promiseFn: init,
  });

  if (isLoading) return options.loading();
  if (error !== undefined) return options.failed(error);

  // проверить на undefined нельзя, потому что init может ничего не вернуть,
  // и это тоже должно считаться успешным завершением
  return options.success(data!);
}

export function useWatch<T>(valueFn: () => T, callback: (newValue: T, oldValue: T) => void, immediately?: boolean) {
  const callbackRef = useRef(callback);
  callbackRef.current = callback;

  const oldValue = valueFn();
  if (immediately) {
    callbackRef.current(oldValue, oldValue);
  }
  useEffect(() => {
    return () => {
      const newValue = valueFn();
      if (newValue === oldValue) {
        return;
      }

      callbackRef.current(newValue, oldValue);
    };
  },        [oldValue]);
}

/**
 * Обертка над useTranslation с возможностью указать путь до перевода
 * @param context {string} где искать перевод
 * @param ns {string} namespace
 * @param options {any} параметры для настройки
 */
export function useTranslate(context: string, ns?: string, options?: any): UseTranslationResponse {
  const def = useTranslation(ns, options);

  return {
    ...def,
    t: (key: string, options?: any) => {
      return def.t(`${context}.${key}`, options);
    }
  }
}
