import axios from "axios";

export const api = axios.create({
  baseURL: "http://localhost:8081", // troca pelo endereço real da API GymFinder
});

