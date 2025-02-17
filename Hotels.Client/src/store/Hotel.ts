import { defineStore } from "pinia";

export type Hotel = {
    _id: string,
    name: string,
    adress: string,
    city: string,
    phone: string,
    description: string,
    HabitacionIds: string[],
    registros: string[],
    created_at: string;
    updated_at: string;

}

export const useHotelsStore = defineStore('hotels', {
    state: () => ({
        hotel: {} as Hotel,
        hotels: [] as Hotel[],
        hotelsUser: [] as Hotel[]

    }),
    actions: {
        GetHotelsOfUser(params: string[]) {
            this.hotelsUser=[];
            for (let i = 0; i < params.length; i++) {
                const id = params[i];
                fetch("/api/hotels/".concat(id))
                    .then((res) => res.json())
                    .then((data) => {
                        data ? this.hotelsUser.push(data) : null;
                    })
                    .catch((err) => {
                        console.log(err);
                    })
            }

        },
        GetHotels() {
            fetch("/api/hotels")
                .then((res) => res.json())
                .then((data) => {
                    this.hotels = data;
                })
                .catch((err) => {
                    console.log(err);
                })
        },
        GetHotel(id: string) {
            fetch("/api/hotels/".concat(id))
                .then((res) => res.json())
                .then((data) => {
                    this.hotel = data;
                    console.log(this.hotel);
                })
                .catch((err) => {
                    console.log(err);
                })
        },
    },
})