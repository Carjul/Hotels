import { jwtDecode } from "jwt-decode";
import { defineStore } from "pinia"
import { ref } from "vue"
import { useRouter } from "vue-router";

export type User = {
    _id: string;
    contraseña: string;
    correo: string;
    imagen: string;
    nombre: string;
    rol: {
        id: number;
        name: string;
    };
    created_at: string;
    updated_at: string;
};

export const useUsuarioStore = defineStore('user', () => {
    const usuario = ref({} as User);
    const router = useRouter();

    function Login(objUser: any) {
        fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(objUser),
        })
            .then((res) => res.json())
            .then((data) => {
                if (data.token) {
                    localStorage.setItem("token", data.token);
                    let dec = jwtDecode(data.token)
                    //@ts-expect-error
                    if (dec.user) {
                        //@ts-expect-error
                        usuario.value = dec.user;
                        localStorage.setItem("user",  JSON.stringify(usuario.value));
                        router.push("/home/dashboard");
                    }
                }

            });
    }

    function LogOut() {
        fetch("/api/logout", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
        })
            .then((res) => res.json())
            .then((data) => {
                console.log(data);
                if (data.message === "Logged out successfully") {
                    usuario.value = {} as User;
                    localStorage.clear();
                    router.push("/");
                }
            });
    }
    return { usuario, Login, LogOut }
})