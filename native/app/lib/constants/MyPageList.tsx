import { Ionicons } from '@expo/vector-icons';
import { MaterialCommunityIcons } from '@expo/vector-icons';
import { FontAwesome } from '@expo/vector-icons';
import { FontAwesome5 } from '@expo/vector-icons';

// マイメニューのリスト
export const myMenuList = [
  {
    title: 'プロフィール',
    icon: <Ionicons name="person" size={24} color="black" />
  },
  {
    title: '友達一覧',
    icon: <MaterialCommunityIcons name="account-group" size={24} color="black" />
  },
  {
    title: 'メッセージボックス',
    icon: <MaterialCommunityIcons name="message" size={24} color="black" />
  },
]

// 読書関連
export const readingRelationList = [
  {
    title: '本棚',
    icon: <MaterialCommunityIcons name="bookshelf" size={24} color="black" />
  },
  {
    title: '自分の感想',
    icon: <MaterialCommunityIcons name="file-document-edit" size={24} color="black" />
  },
  {
    title: '新刊チェック',
    icon: <FontAwesome5 name="calendar-alt" size={24} color="black" />
  }
]

// フリマ関連
export const fleaMarketRelationList = [
  {
    title: '出品リスト',
    icon: <MaterialCommunityIcons name="book-plus-multiple" size={24} color="black" />
  },
  {
    title: '購入リスト',
    icon: <MaterialCommunityIcons name="cart" size={24} color="black" />
  },
  {
    title: '売り上げ・申請',
    icon: <MaterialCommunityIcons name="piggy-bank" size={24} color="black" />
  },
  {
    title: '最近見た商品',
    icon: <FontAwesome5 name="history" size={24} color="black" />
  },
  {
    title: 'コメントした商品',
    icon: <FontAwesome name="comments" size={24} color="black" />
  }
]

// その他
export const otherList = [
  {
    title: 'お知らせ',
  },
  {
    title: 'お問い合わせ',
  },
  {
    title: 'ヘルプ',
  },
  {
    title: '設定',
  }
]
