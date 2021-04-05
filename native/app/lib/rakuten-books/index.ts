import axios from 'axios';

const baseUrl = 'https://app.rakuten.co.jp/services/api/BooksBook/Search';
const version = '20170404';
const format = 'json';
const formatVersion = 2;
const applicationId = '1053511986137691200';


export async function searchBook(param: string) {
  const url = `${baseUrl}/${version}?format=${format}&title=${encodeURI(param)}&formatVersion=${formatVersion}&applicationId=${applicationId}`;

  return axios.get(url)
    .then((res) => {
      console.log(res.data);
      return res;
    })
    .catch((err) => {
      return err;
    });
}
