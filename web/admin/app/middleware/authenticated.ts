const excludedPaths: string[] = ['/', '/signin']

export default async ({ route, store, redirect }) => {
  if (excludedPaths.includes(route.path)) {
    return
  }

  await store.dispatch('auth/authorization').catch(() => {
    redirect('/signin')
  })
}
