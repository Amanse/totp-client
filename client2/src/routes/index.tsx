import { component$ } from "@builder.io/qwik";
import axios from "axios";
import type { DocumentHead } from "@builder.io/qwik-city";

axios.defaults.baseURL = "http://localhost:8080"

export default component$(() => {
  return (
    <>
    </>
  );
});

export const head: DocumentHead = {
  title: "Welcome to Qwik",
  meta: [
    {
      name: "description",
      content: "Qwik site description",
    },
  ],
};
