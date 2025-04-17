static class LogLine
{
    public static string Message(string logLine)
    {
        return logLine.Substring(logLine.IndexOf(":") + 1).Trim();
    }

    public static string LogLevel(string logLine)
    {
        var startIndex = logLine.IndexOf('[') + 1;
        var length = logLine.IndexOf(']') - startIndex;
        return logLine.Substring(startIndex, length).Trim().ToLower();
    }

    public static string Reformat(string logLine)
    {
        return $"{Message(logLine)} ({LogLevel(logLine)})";
    }
}
