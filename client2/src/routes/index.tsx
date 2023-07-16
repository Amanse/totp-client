import { component$, useSignal, useStore, useVisibleTask$ } from "@builder.io/qwik";
import type { Signal } from "@builder.io/qwik";
import axios from "axios";
import type { DocumentHead } from "@builder.io/qwik-city";
import { useNavigate } from "@builder.io/qwik-city";

axios.defaults.baseURL = "http://localhost:8080"

export default component$(() => {
    const nav = useNavigate()
    const token: Signal<string | null> = useSignal<string | null>(null)
    const mainStuff = useStore({ "issuers": {} })
    useVisibleTask$(() => {
        if (!localStorage.getItem("token")) {
            nav("/signup")
        } else {
            token.value = localStorage.getItem("token")
            axios.get("http://localhost:8080/codes", {
                headers: {
                    "Authorization": ` Bearer ${token.value}`
                }
            }).then(d => {
                mainStuff.issuers = d.data.issuers;
                console.log(mainStuff.issuers)
            })
        }

    })
    return (
        <>
        </>
    );
});

export const head: DocumentHead = {
    title: "Welcome to totp",
    meta: [
        {
            name: "description",
            content: "Qwik site description",
        },
    ],
};
