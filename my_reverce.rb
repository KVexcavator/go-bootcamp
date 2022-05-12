class Array
  def my_reverce
    res,i,l = [],-1,size
    while l >= 1
      res << self[i]
      l -= 1
      i -= 1
    end
    res
  end
end

c = ["a","b","c"].my_reverce
p c
