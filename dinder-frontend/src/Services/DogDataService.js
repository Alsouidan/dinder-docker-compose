import axios from 'axios'

const GO_API_URL = 'https://dinder-backend.herokuapp.com';
const DOGS_API_URL = `${GO_API_URL}/dogs`

class DogDataService {
    async retrieveAllDogs()
    {
        try{
            const response=await  axios.get(`${DOGS_API_URL}`,);
            return response.data
        }
        catch(error){}    
    }
    async retrieveDog(id)
    {
        try{
            console.log(id)
            const response= await axios.get(`${DOGS_API_URL}/${id}`);
            return response.data
        }
        catch(error){}
    }
    async createDog(dog)
    {
        try{
            const response=await  axios.post(`${DOGS_API_URL}`,dog);
            return response.data
        }
        catch(error){}
    }
    async deleteDog(id){
        try{
            const response=await  axios.delete(`${DOGS_API_URL}/${id}`);
            return response.data
         }
        catch(error){}
        }
    async updateDog(id,dog){
        try{
            const response=await  axios.put(`${DOGS_API_URL}/${id}`,dog);
            return response.data
        }
        catch(error){}
    }
    async approveDog(id1,id2){
        try{
            const response=await  axios.put(`${DOGS_API_URL}/approve/${id1}/${id2}`);
            return response.data
        }
        catch(error){}
    }
    async rejectDog(id1,id2){
        try{
            const response=await  axios.put(`${DOGS_API_URL}/reject/${id1}/${id2}`);
            return response.data
        }
        catch(error){}
    }
    async recomendDogs(specs){
        try{
            console.log(specs)
            const response=await  axios.post(`${DOGS_API_URL}/getRec`,specs);
            return response.data
        }
        catch(error){}
    }
    async getDogMatches(id){
        try{
            const response=await  axios.get(`${DOGS_API_URL}/getMatches/${id}`);
            return response.data
        }
        catch(error){}
    }
}
    

export default new DogDataService();