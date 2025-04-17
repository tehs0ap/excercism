public static class LogAnalysis 
{
    
    public static string SubstringAfter(this string str, string delimiter)
    {
        if (delimiter.Length == 0 || str.Length == 0)
        {
            return str;
        }
        
        var delimiterIndex = str.IndexOf(delimiter, StringComparison.Ordinal);
        if (delimiterIndex < 0)
        {
            return str;
        }
        
        return  str.Substring(delimiterIndex + delimiter.Length);
    }

    public static string SubstringBetween(this string str, string delimiter1, string delimiter2)
    {
        if (str.Length == 0 || delimiter1.Length == 0 || delimiter2.Length == 0)
        {
            return str;
        }
        
        var delimiter1Index = str.IndexOf(delimiter1, StringComparison.Ordinal);
        var delimiter2Index = str.IndexOf(delimiter2, StringComparison.Ordinal);
        if (delimiter1Index < 0 || delimiter2Index < 0)
        {
            return str;
        }
        
        var startIndex = delimiter1Index + delimiter1.Length;
        var length = delimiter2Index - startIndex;
        return  str.Substring(startIndex, length);
    }
    
    public static string Message(this string str)
    {
        return str.SubstringAfter(":").Trim();
    }

    public static string LogLevel(this string str)
    {
        return str.SubstringBetween("[", "]").Trim();        
    }
}