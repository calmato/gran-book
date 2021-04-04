import firebase from 'firebase';

export type Message = {
    text: string;
    createdAt: firebase.firestore.Timestamp;
    userId: string;
}