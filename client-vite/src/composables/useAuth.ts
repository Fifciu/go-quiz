import axios from '../axios';

const useAuth = () => {
  const login = async (email: string, password: string): Promise<any> => {
    try {
      const response = await axios.post('/auth', {
        email,
        password
      });
    } catch (err) {
      console.log('Error at useAuth/login: ', err);
    }
  }

  return {
    login
  }
};

export { useAuth };
