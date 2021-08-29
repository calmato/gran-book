import { useCallback, useEffect, useState } from 'react';
import { getRecommendBooks } from '~/lib/rakuten-books';
import { ISearchResultItem } from '~/types/response/external/rakuten-books';

/**
 * おすすめ書籍を取得するカスタムhooks
 * @returns
 */
export function useRecommendBooks() {
  const [recommendBooks, setRecommendBooks] = useState<ISearchResultItem[]>([]);

  const [_isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<boolean>(true);

  const fetch = useCallback(async () => {
    setIsLoading(true);
    try {
      const res = await getRecommendBooks();
      setRecommendBooks(res.Items);
    } catch (e) {
      setError(true);
    }
    setIsLoading(false);
  }, [setIsLoading, setRecommendBooks, setError]);

  useEffect(() => {
    fetch();
  }, []);

  return { recommendBooks, error, fetch };
}
