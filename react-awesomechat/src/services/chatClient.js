import { Empty, User, SentMessage, Group } from '../contracts/awesomechat_pb.js';
import { ChatCoreClient } from '../contracts/awesomechat_grpc_web_pb.js';

const proxyHost = 'http://' + process.env.REACT_APP_PROXY_HOST + ':8080'
const client = new ChatCoreClient(proxyHost, null, null)

let token = ""

export async function createUser(username) {
  var request = new User();
  request.setUsername(username)

  return new Promise((resolve, reject) => {
    client.createUser(request, null, (err, response) => {
      if (err)
        return reject(err)

      token = response.toObject().token
      console.log(token);
      resolve(token)
    })
  })
}

export async function connectStream(username) {
  const request = new User()
  const metadata = { 'authorization': token }

  request.setUsername(username)

  const chatStream = await client.connect(request, metadata)
  return chatStream
}

export async function sendMessage(content, receiver) {
  const request = new SentMessage();
  const metadata = { 'authorization': token }

  request.setReceiver(receiver)
  request.setContent(content)

  return new Promise((resolve, reject) => {
    client.sendMessage(request, metadata, (err, resp) => {
      if (err)
        return reject(err)

      resolve(resp)
    })
  })
}

export async function createGroupChat(groupName) {
  const request = new Group();
  const metadata = { 'authorization': token }

  request.setGroupname(groupName)

  return new Promise((resolve, reject) => {
    client.createGroupChat(request, metadata, (err, resp) => {
      if (err)
        return reject(err)

      resolve(resp)
    })
  })
}

export async function joinGroupChat(groupName) {
  const request = new Group();
  const metadata = { 'authorization': token }

  request.setGroupname(groupName)

  return new Promise((resolve, reject) => {
    client.joinGroupChat(request, metadata, (err, resp) => {
      if (err)
        return reject(err)

      resolve(resp)
    })
  })
}

export async function leftGroupChat(groupName) {
  const request = new Group();
  const metadata = { 'authorization': token }

  request.setGroupname(groupName)

  return new Promise((resolve, reject) => {
    client.leftGroupChat(request, metadata, (err, resp) => {
      if (err)
        return reject(err)

      resolve(resp)
    })
  })
}

export async function listChannels() {
  const request = new Empty();
  const metadata = { 'authorization': token }

  return new Promise((resolve, reject) => {
    client.listChannels(request, metadata, (err, resp) => {
      if (err)
        return reject(err)

      const result = resp.toObject()
      resolve(result)
    })
  })
}

