import http from '@/store/helpers/http';

const putUser = async (data) => http().put(`/user`, {
  username: data.username,
  email: data.email,
  password: data.password,
});

export { putUser as default };
