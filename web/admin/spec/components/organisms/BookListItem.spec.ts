import { mount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import BookListItem from '~/components/organisms/BookListItem.vue'

describe('components/organisms/BookListItem', () => {
  let wrapper: any

  beforeEach(() => {
    wrapper = mount(BookListItem, { ...Options })
  })

  describe('script', () => {
    describe('props', () => {
      describe('book', () => {
        it('初期値', () => {
          expect(wrapper.props().book).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({
            book: {
              id: 1,
              title: '小説　ちはやふる　上の句',
              titleKana: 'ショウセツ チハヤフルカミノク',
              description:
                '綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。',
              isbn: '9784062938426',
              publisher: '講談社',
              publishedOn: '2018年01月16日',
              thumbnailUrl:
                'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120',
              rakutenUrl: 'https://books.rakuten.co.jp/rb/15271426/',
              rakutenGenreId: '001004008001/001004008003/001019001',
              author: '末次 由紀',
              authorKana: 'スエツグ ユキ',
              createdAt: '2021-05-02 14:38:27',
              updatedAt: '2021-05-02 14:38:27',
            },
          })
          expect(wrapper.props().book).toEqual({
            id: 1,
            title: '小説　ちはやふる　上の句',
            titleKana: 'ショウセツ チハヤフルカミノク',
            description:
              '綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。',
            isbn: '9784062938426',
            publisher: '講談社',
            publishedOn: '2018年01月16日',
            thumbnailUrl:
              'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120',
            rakutenUrl: 'https://books.rakuten.co.jp/rb/15271426/',
            rakutenGenreId: '001004008001/001004008003/001019001',
            author: '末次 由紀',
            authorKana: 'スエツグ ユキ',
            createdAt: '2021-05-02 14:38:27',
            updatedAt: '2021-05-02 14:38:27',
          })
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClick()
          expect(wrapper.emitted('click')).toBeTruthy()
        })
      })
    })
  })
})
