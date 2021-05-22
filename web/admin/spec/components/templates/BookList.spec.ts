import { mount } from '@vue/test-utils'
import * as Options from '~~/spec/helpers/component-helper'
import BookList from '~/components/templates/BookList.vue'
import { BookSearchOptions, IBookSearchForm, IBookSearchParams } from '~/types/forms'
import { IBook } from '~/types/store'

describe('components/templates/BookList', () => {
  let wrapper: any

  beforeEach(() => {
    const params: IBookSearchParams = {
      title: '',
      author: '',
      publisher: '',
      isbn: '',
      size: 2,
    }
    const searchForm: IBookSearchForm = { params, options: BookSearchOptions }

    const books: IBook[] = [
      {
        id: 1,
        title: '小説　ちはやふる　上の句',
        titleKana: 'ショウセツ チハヤフルカミノク',
        description:
          '綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。',
        isbn: '9784062938426',
        publisher: '講談社',
        publishedOn: '2018年01月16日',
        thumbnailUrl: 'https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120',
        rakutenUrl: 'https://books.rakuten.co.jp/rb/15271426/',
        rakutenGenreId: '001004008001/001004008003/001019001',
        author: '末次 由紀',
        authorKana: 'スエツグ ユキ',
        createdAt: '2021-05-02 14:38:27',
        updatedAt: '2021-05-02 14:38:27',
      },
    ]

    wrapper = mount(BookList, {
      ...Options,
      propsData: { books, searchForm },
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('books', () => {
        it('初期値', () => {
          expect(wrapper.props().books).toEqual([
            {
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
          ])
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ books: [] })
          expect(wrapper.props().books).toEqual([])
        })
      })
    })

    describe('methods', () => {
      describe('onSubmitSearchForm', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitSearchForm()
          expect(wrapper.emitted('submit')).toBeTruthy()
        })
      })

      describe('onClickBookCard', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickBookCard({
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
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toEqual({
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

      describe('onClearSearchForm', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClearSearchForm()
          expect(wrapper.emitted('clear')).toBeTruthy()
        })
      })
    })
  })
})
