angular.module('app.auth', [])



.factory('authFactory', function() {
  return {
    getToken: function() {
      return Cookies.getJSON('auth').token;
    },
    getUsername: function() {
      return Cookies.getJSON('auth').username;
    },
    getGravatar: function() {
      return Cookies.getJSON('auth').gravatar;
    },
    getAuth: function() {
      return Cookies.getJSON('auth');
    },
    setAuth: function(token, username, gravatar, expire) {
      Cookies.set('auth', JSON.stringify({token:token, username:username, gravatar:gravatar}), {expires: expire});
    },
    killAuth: function() {
      Cookies.remove('auth');
    }
  }
});
