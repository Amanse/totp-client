import { component$, useSignal, $ } from '@builder.io/qwik';
import { useNavigate } from '@builder.io/qwik-city';
import axios from 'axios';
import { Client } from '@passwordlessdev/passwordless-client';

export default component$(() => {
    let username = useSignal("");
    let password = useSignal("");
    let errorString = useSignal("");
    let ifError = useSignal(false);

    const nav = useNavigate();

    const signUpPasskey = $(async () => {
        const p = new Client({
            apiKey: "totp:public:4f46665f2e6a41f3abdbdae8ef539114"
        });

        axios.post("http://localhost:8080/signup?isPasskey=1", { username: username.value, password: "hehe" }).then(r => {
            p.register(r.data.token).then(({ token, error }) => {
                if (token) {
                    localStorage.setItem("token", token)
                }
            });

        })

    })


    const signUp = $(() => {
        username.value = username.value.trim()
        password.value = password.value.trim()
        try {
            axios.post("http://localhost:8080/signup", {
                username: username.value,
                password: password.value
            }).then(data => {
                localStorage.setItem("token", data.data.token)
                ifError.value = false
                errorString.value = ""
                nav("/login")
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
            <button type="submit" onClick$={() => signUp()}>Signup</button>
            <button type="submit" onClick$={() => signUpPasskey()}>Signup with passkey</button>
            {ifError &&
                <>
                    {errorString}
                </>
            }
        </>
    )
})
