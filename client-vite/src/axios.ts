import axios from 'axios';

const instance = axios.create({
  baseURL: `${process.env.api_protocol}://${process.env.api_host}:${process.env.api_port}`
});

export default instance;
