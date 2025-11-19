import axios from "axios";

export const api = axios.create({
  baseURL: "https://gymfinder-m411.onrender.com", // troca pelo endereço real da API GymFinder
});

