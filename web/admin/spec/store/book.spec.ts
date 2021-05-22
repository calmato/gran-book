import { setup, refresh } from '~~/spec/helpers/store-helper'
import { BookStore } from '~/store'

describe('store/book', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getBooks', () => {
      expect(BookStore.getBooks).toEqual([])
    })
  })

  // Jestの設定周り整えないと行けないので後ほど作成
  // describe('actions', () => {
  //   describe('searchBookFromRakutenBooksAPI', () => {})
  // })
})
