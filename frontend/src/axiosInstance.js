import axios from "axios";

const client = axios.create({
  baseURL: process.env.REACT_APP_GO,
  headers: {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
  timeout: 60 * 1000,
});

export default client;
