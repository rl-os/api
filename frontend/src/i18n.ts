import Backend from "i18next-xhr-backend";
import i18next from "i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";
import Config from './config';

export const SUPPORTED_LANGUAGES = [
  {code: "ru", name: "Русский"},
  {code: "en", name: "English"},
];

const i18n = i18next
  .use(Backend)
  .use(LanguageDetector)
  .use(initReactI18next);

// for all options read
// https://www.i18next.com/overview/configuration-options
i18n.init({
  debug: true,
  react: {
    wait: true,
    useSuspense: true,
  },

  fallbackLng: Config.i18n.fallbackLng,
  whitelist: SUPPORTED_LANGUAGES.map(lang => lang.code),

  interpolation: {
    escapeValue: false, // not needed for react as it escapes by default
  },
  backend: {
    loadPath: "/locales/{{lng}}.json",
  },
});

export default i18n;
