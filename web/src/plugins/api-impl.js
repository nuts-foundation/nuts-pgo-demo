export default createApi;
function createApi(options) {
  const basePath = '/web';
  const endpoint = options.endpoint || '';
  const cors = !!options.cors;
  const mode = cors ? 'cors' : 'basic';
  const buildQuery = (obj) => {
    return Object.keys(obj)
      .filter(key => typeof obj[key] !== 'undefined')
      .map((key) => {
        const value = obj[key];
        if (value === undefined) {
          return '';
        }
        if (value === null) {
          return key;
        }
        if (Array.isArray(value)) {
          if (value.length) {
            return key + '=' + value.map(encodeURIComponent).join('&' + key + '=');
          } else {
            return '';
          }
        } else {
          return key + '=' + encodeURIComponent(value);
        }
      }).join('&');
    };
  return {
    createAccount(parameters) {
      const params = typeof parameters === 'undefined' ? {} : parameters;
      let headers = {
        'content-type': 'application/json',

      };
      return fetch(endpoint + basePath + '/account'
        , {
          method: 'POST',
          headers,
          mode,
          body: JSON.stringify(params['body']),

        });
    },
    createSession(parameters) {
      const params = typeof parameters === 'undefined' ? {} : parameters;
      let headers = {
        'content-type': 'application/json',

      };
      return fetch(endpoint + basePath + '/auth'
        , {
          method: 'POST',
          headers,
          mode,
          body: JSON.stringify(params['body']),

        });
    },
    checkSession(parameters) {
      const params = typeof parameters === 'undefined' ? {} : parameters;
      let headers = {

      };
      return fetch(endpoint + basePath + '/private'
        , {
          method: 'GET',
          headers,
          mode,
        });
    },

  };
}
