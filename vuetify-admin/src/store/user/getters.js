import { Base64 } from 'js-base64';

export default {
  userInfo(state) {
    if (state.token) {
      const jwtTokens = state.token.split('.');
      if (jwtTokens.length > 1) {
        return JSON.parse(Base64.decode(jwtTokens[1]));
      }
    }
    return '';
  },
};
