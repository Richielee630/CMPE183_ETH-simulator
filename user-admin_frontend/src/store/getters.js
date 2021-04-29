const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  userid: state => state.user.userid,
  usertype: state => state.user.usertype,
  name: state => state.user.name
}
export default getters
