import { setup, refresh } from '~~/spec/helpers/store-helper'
import { CommonStore } from '~/store'
import { ISnackbar } from '~/types/store'
import { ApiError } from '~/types/exception'

describe('store/common', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getSnackbarColor', () => {
      expect(CommonStore.getSnackbarColor).toBe('info')
    })

    it('getSnackbarMessage', () => {
      expect(CommonStore.getSnackbarMessage).toBe('')
    })
  })

  describe('actions', () => {
    describe('showSnackbar', () => {
      it('stateが更新されること', () => {
        const payload: ISnackbar = { color: 'primary', message: 'テストメッセージ' }
        CommonStore.showSnackbar(payload)
        expect(CommonStore.getSnackbarColor).toBe('primary')
        expect(CommonStore.getSnackbarMessage).toBe('テストメッセージ')
      })
    })

    describe('showSuccessInSnackbar', () => {
      it('stateが更新されること', () => {
        CommonStore.showSuccessInSnackbar('テストメッセージ')
        expect(CommonStore.getSnackbarColor).toBe('success')
        expect(CommonStore.getSnackbarMessage).toBe('テストメッセージ')
      })
    })

    describe('showErrorInSnackbar', () => {
      describe('引数がApiError型のとき', () => {
        it('ステータスコードが400のとき', () => {
          const err = new ApiError(400, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe('入力内容に誤りがあります。')
        })

        it('ステータスコードが401のとき', () => {
          const err = new ApiError(401, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe('認証エラーです。再度ログインしてください。')
        })

        it('ステータスコードが403のとき', () => {
          const err = new ApiError(403, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe('権限エラーです。')
        })

        it('ステータスコードが404のとき', () => {
          const err = new ApiError(404, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe(
            '正常に処理できませんでした。しばらく経ってからもう一度お試しください。'
          )
        })

        it('ステータスコードが409のとき', () => {
          const err = new ApiError(409, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe('既に登録されています。')
        })

        it('ステータスコードが500のとき', () => {
          const err = new ApiError(500, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe(
            'サーバーエラーが発生しました。しばらく経ってからもう一度お試しください。'
          )
        })

        it('ステータスコードが501のとき', () => {
          const err = new ApiError(501, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe(
            'サーバーエラーが発生しました。しばらく経ってからもう一度お試しください。'
          )
        })

        it('ステータスコードが503のとき', () => {
          const err = new ApiError(503, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe(
            'サーバーエラーが発生しました。しばらく経ってからもう一度お試しください。'
          )
        })

        it('ステータスコードが504のとき', () => {
          const err = new ApiError(504, 'api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe('通信がタイムアウトしました。再度実行してください。')
        })

        it('ステータスコードがその他のとき', () => {
          const err = new ApiError(301, 'redirect')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe(
            '予期せぬエラーが発生しました。しばらく経ってからもう一度お試しください。'
          )
        })
      })

      describe('引数がその他Error型のとき', () => {
        it('stateが更新されること', () => {
          const err = new Error('api error')
          CommonStore.showErrorInSnackbar(err)
          expect(CommonStore.getSnackbarColor).toBe('error')
          expect(CommonStore.getSnackbarMessage).toBe(
            '予期せぬエラーが発生しました。しばらく経ってからもう一度お試しください。'
          )
        })
      })
    })

    describe('hiddenSnackbar', () => {
      beforeEach(() => {
        const payload: ISnackbar = { color: 'primary', message: 'テストメッセージ' }
        CommonStore.showSnackbar(payload)
      })

      it('stateが初期値に戻ること', () => {
        CommonStore.hiddenSnackbar()
        expect(CommonStore.getSnackbarColor).toBe('info')
        expect(CommonStore.getSnackbarMessage).toBe('')
      })
    })
  })
})
