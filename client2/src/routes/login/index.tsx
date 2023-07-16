import { component$, useSignal, $ } from "@builder.io/qwik"
import { useNavigate } from "@builder.io/qwik-city";
import { Client } from '@passwordlessdev/passwordless-client';
import axios from "axios";
import { sha256 } from "js-sha256";

export default component$(() => {
    let username = useSignal("");
    let password = useSignal("");
    let errorString = useSignal("");
    let ifError = useSignal(false);

    const nav = useNavigate();

    const loginWithPasskey = $(() => {
        const p = new Client({
            apiKey: "totp:public:4f46665f2e6a41f3abdbdae8ef539114"
        });

        p.signinWithId(username.value).then(({ token, error }) => {
            if (error) {
                console.log(error)
                return
            }

            axios.post("http://localhost:8080/login/passkey", { token }).then(data => {
                localStorage.setItem("token", data.data.token)
                ifError.value = false
                errorString.value = ""
                nav("/")
            })
        });

    })


    const login = $(() => {
        try {
            axios.post("http://localhost:8080/login", {
                username: username.value.trim(),
                password: sha256(password.value.trim())
            }).then(data => {
                localStorage.setItem("token", data.data.token)
                ifError.value = false
                errorString.value = ""
                nav("/")
            }).catch(err => {
                if (err.response.status == 400) {
                    errorString.value = err.response.data
                    ifError.value = true;
                }
            })
        } catch (error) {
            console.log(error)
        }

    })
    return (
        <>
            <input type="text" bind: value={username} placeholder='username' onInput$={(e) => (username.value = (e.target as HTMLInputElement).value)} />
            <input type="text" bind: value={password} placeholder='password' onInput$={(e) => (password.value = (e.target as HTMLInputElement).value)} />
            <button type="submit" onClick$={() => login()}>login</button>
            <button type="submit" onClick$={() => loginWithPasskey()}>login with passkey</button>
            {ifError &&
                <>
                    {errorString}
                </>
            }
        </>
    )
})
