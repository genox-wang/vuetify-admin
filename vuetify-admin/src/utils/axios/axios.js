import axios from 'axios';
import config from '@/config';


axios.defaults.baseURL = config.apiBaseURL;
axios.defaults.transformRequest = [data => JSON.stringify(data)];
// axios.defaults.withCredentials = true;
// axios.defaults.validateStatus = status => status >= 200 && status < 500;

export default axios;

