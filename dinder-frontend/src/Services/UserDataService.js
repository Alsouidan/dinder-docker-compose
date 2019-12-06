import axios from 'axios'

const GO_API_URL = 'https://dinder-backend.herokuapp.com';
const USER_API_URL = `${GO_API_URL}/users`
//const TOKEN =''

class  UserDataService {
    async retrieveAllUsers() {
        try{
             const response=await axios.get(`${USER_API_URL}`)
             console.log(response)
        return response.data
        //{ headers: { authorization: 'Basic ' + window.btoa(USER + ":" + Password) }} 
        }
        catch(error){}
    }
    async retrieveUser(id) {
        //console.log('executed service')
        try{
        const response = await axios.get(`${USER_API_URL}/${id}`);
        return response.data
    }
    catch(error){}
    }
    async createUser(user) {
        try{

        const response = await axios.post(`${USER_API_URL}`,user);
        return response.data
        }
    catch(error){}
    }
    async deleteUser(id) {
        try{
        //console.log('executed service')
        const response = await axios.delete(`${USER_API_URL}/${id}`);
        return response.data
    }
    catch(error){}
    }

    async updateUser(id, user) {
        try {
        //console.log('executed service')
        const response = await axios.put(`${USER_API_URL}/${id}`, user);
        return response.data
    }
    catch(error){}
    }
    async addDog(id,_id){
        try{
            const response = await axios.put(`${USER_API_URL}/addDog/${id}`,_id);
            return response.data
        }
        catch(error){}
    }
    async retrieveDogs(id) {
        //console.log('executed service')
        try{
        const response = await axios.get(`${USER_API_URL}/getDogs/${id}`);
        return response.data
    }
    catch(error){}
    }
    login(creds) {
        try{
            console.log("Here");
        //console.log('executed service')
        return axios.post(`${USER_API_URL}/login`,creds);
        
    }
    catch(error){}
    }
    authUser(token){
        axios.defaults.headers.common['Authorization'] = token;
    }


}

export default new UserDataService();