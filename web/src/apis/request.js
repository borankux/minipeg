import axios from 'axios'

let baseUrl = process.env.NODE_ENV === 'development'?
  'http://localhost:8033':
  'http://106.14.35.196:8033';

const request = axios.create({
  baseURL : baseUrl
});

export default request
