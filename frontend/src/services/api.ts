import axios from "axios";

export const api = axios.create({
  baseURL: "https://gymfinder-1.onrender.com", // troca pelo endereço real da API GymFinder
});

