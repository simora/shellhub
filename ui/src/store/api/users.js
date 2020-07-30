import http from '@/store/helpers/http';

const putUser = async (data) => http().patch(`/user`, {
  username: data.username,
  email: data.email,
  password: data.password,
});

export { putUser as default };
