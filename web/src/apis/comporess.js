import request from './request'

export function uploadFile(file_name, content)
{
  return request({
    method:'post',
    url:'/compress',
    data: {
      file_name:file_name,
      content: content
    }
  });
}
