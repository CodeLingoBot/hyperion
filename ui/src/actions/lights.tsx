import apiFetch from './index';

export const RECEIVE_LIGHT_LIST = 'RECEIVE_LIGHT_LIST';
export const RECEIVE_STATE_LIST = 'RECEIVE_STATE_LIST';

export function fetchLightList() {
  return (dispatch: any) => {
    return apiFetch('lights')
      .then(response => response.json())
      .then(json => dispatch(receiveLightList(json)));
  };
}
function receiveLightList(lights: any) {
  return {
    type: RECEIVE_LIGHT_LIST,
    lights
  };
}

export function sendCommands(commands: string[]) {
  return (dispatch: any) => {
    return apiFetch('commands', {
      method: 'POST',
      body: JSON.stringify(commands)
    }).then(response => response.json());
  };
}
